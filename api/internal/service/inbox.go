package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jhillyerd/enmime"
	"github.com/microcosm-cc/bluemonday"
	"ivpn.net/email/api/internal/model"
)

var (
	ErrGetInboxMessages   = errors.New("Unable to retrieve inbox messages.")
	ErrGetInboxUnread     = errors.New("Unable to retrieve unread count.")
	ErrGetInboxMessage    = errors.New("Unable to retrieve inbox message.")
	ErrDeleteInboxMessage = errors.New("Unable to delete inbox message.")
	ErrInboxMessageSize   = errors.New("inbox message exceeds size limit")
)

type InboxStore interface {
	PostInboxMessage(context.Context, model.InboxMessage) error
	GetInboxMessages(context.Context, string, string) ([]model.InboxMessage, error)
	GetInboxUnreadCount(context.Context, string) (int64, error)
	GetInboxMessage(context.Context, uint, string) (model.InboxMessage, error)
	DeleteInboxMessage(context.Context, uint, string) error
}

// inboxHTMLPolicy sanitizes email HTML for browser display: no scripts, no
// styles, no images (blocks remote-content tracking), links forced to
// nofollow + target=_blank with rel=noopener.
var inboxHTMLPolicy = func() *bluemonday.Policy {
	p := bluemonday.NewPolicy()
	p.AllowStandardURLs()
	p.AllowStandardAttributes()
	p.AllowLists()
	p.AllowTables()
	p.AllowElements(
		"a", "b", "blockquote", "br", "caption", "code", "del", "div", "em",
		"h1", "h2", "h3", "h4", "h5", "h6", "hr", "i", "ins", "p", "pre", "q",
		"s", "small", "span", "strike", "strong", "sub", "sup", "u",
	)
	p.AllowAttrs("href").OnElements("a")
	p.RequireNoFollowOnLinks(true)
	p.AddTargetBlankToFullyQualifiedLinks(true)
	return p
}()

// StoreInboxMessage persists an incoming email for an inbox-type alias and
// records an Inbox stats row.
func (s *Service) StoreInboxMessage(ctx context.Context, alias model.Alias, msg model.Msg, raw []byte) error {
	if len(raw) > model.MaxInboxMessageSize {
		return fmt.Errorf("%w: %d bytes", ErrInboxMessageSize, len(raw))
	}

	err := s.Store.PostInboxMessage(ctx, model.InboxMessage{
		UserID:   alias.UserID,
		AliasID:  alias.ID,
		From:     msg.From,
		FromName: msg.FromName,
		Subject:  msg.Subject,
		Size:     len(raw),
		Raw:      raw,
	})
	if err != nil {
		return err
	}

	if err := s.SaveMessage(ctx, alias, model.Inbox); err != nil {
		log.Println("error saving inbox stats message", err)
	}

	return nil
}

func (s *Service) GetInboxMessages(ctx context.Context, aliasID string, userID string) ([]model.InboxMessage, error) {
	messages, err := s.Store.GetInboxMessages(ctx, aliasID, userID)
	if err != nil {
		log.Printf("error getting inbox messages: %s", err.Error())
		return nil, ErrGetInboxMessages
	}

	return messages, nil
}

func (s *Service) GetInboxUnreadCount(ctx context.Context, userID string) (int64, error) {
	count, err := s.Store.GetInboxUnreadCount(ctx, userID)
	if err != nil {
		log.Printf("error getting inbox unread count: %s", err.Error())
		return 0, ErrGetInboxUnread
	}

	return count, nil
}

// GetRenderedInboxMessage loads a stored message, parses its MIME structure
// and returns a sanitized representation safe for browser display.
func (s *Service) GetRenderedInboxMessage(ctx context.Context, ID uint, userID string) (model.RenderedMessage, error) {
	message, err := s.Store.GetInboxMessage(ctx, ID, userID)
	if err != nil {
		log.Printf("error getting inbox message: %s", err.Error())
		return model.RenderedMessage{}, ErrGetInboxMessage
	}

	return renderInboxMessage(message)
}

func (s *Service) DeleteInboxMessage(ctx context.Context, ID uint, userID string) error {
	err := s.Store.DeleteInboxMessage(ctx, ID, userID)
	if err != nil {
		log.Printf("error deleting inbox message: %s", err.Error())
		return ErrDeleteInboxMessage
	}

	return nil
}

func renderInboxMessage(message model.InboxMessage) (model.RenderedMessage, error) {
	env, err := enmime.ReadEnvelope(bytes.NewReader(message.Raw))
	if err != nil {
		log.Printf("error parsing inbox message MIME: %s", err.Error())
		return model.RenderedMessage{}, ErrGetInboxMessage
	}

	attachments := make([]model.Attachment, 0, len(env.Attachments))
	for _, a := range env.Attachments {
		attachments = append(attachments, model.Attachment{
			Name: a.FileName,
			Size: len(a.Content),
		})
	}

	return model.RenderedMessage{
		ID:          message.ID,
		CreatedAt:   message.CreatedAt,
		From:        message.From,
		FromName:    message.FromName,
		Subject:     message.Subject,
		HTML:        inboxHTMLPolicy.Sanitize(env.HTML),
		Text:        env.Text,
		Attachments: attachments,
	}, nil
}
