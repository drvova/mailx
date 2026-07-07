package service

import (
	"strings"
	"testing"

	"ivpn.net/email/api/internal/model"
)

func rawMail(html string) []byte {
	return []byte("From: attacker@evil.test\r\n" +
		"To: victim@inbox.test\r\n" +
		"Subject: payload\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=utf-8\r\n" +
		"\r\n" +
		html + "\r\n")
}

func render(t *testing.T, html string) model.RenderedMessage {
	t.Helper()
	rendered, err := renderInboxMessage(model.InboxMessage{Raw: rawMail(html)})
	if err != nil {
		t.Fatalf("renderInboxMessage failed: %v", err)
	}
	return rendered
}

func TestRenderInboxMessageStripsScripts(t *testing.T) {
	rendered := render(t, `<p>hi</p><script>alert(1)</script>`)
	if strings.Contains(rendered.HTML, "<script") || strings.Contains(rendered.HTML, "alert(1)") {
		t.Fatalf("script not stripped: %q", rendered.HTML)
	}
	if !strings.Contains(rendered.HTML, "<p>hi</p>") {
		t.Fatalf("benign content lost: %q", rendered.HTML)
	}
}

func TestRenderInboxMessageStripsEventHandlers(t *testing.T) {
	rendered := render(t, `<p onmouseover="steal()">text</p><div onclick="x()">y</div>`)
	if strings.Contains(rendered.HTML, "onmouseover") || strings.Contains(rendered.HTML, "onclick") {
		t.Fatalf("event handler not stripped: %q", rendered.HTML)
	}
}

func TestRenderInboxMessageStripsRemoteImages(t *testing.T) {
	rendered := render(t, `<p>a</p><img src="https://tracker.evil.test/pixel.gif">`)
	if strings.Contains(rendered.HTML, "<img") || strings.Contains(rendered.HTML, "tracker.evil.test") {
		t.Fatalf("remote image not stripped: %q", rendered.HTML)
	}
}

func TestRenderInboxMessageStripsJavascriptLinks(t *testing.T) {
	rendered := render(t, `<a href="javascript:alert(1)">click</a>`)
	if strings.Contains(rendered.HTML, "javascript:") {
		t.Fatalf("javascript: URL not stripped: %q", rendered.HTML)
	}
}

func TestRenderInboxMessageHardensLinks(t *testing.T) {
	rendered := render(t, `<a href="https://example.test/x">link</a>`)
	if !strings.Contains(rendered.HTML, `rel=`) || !strings.Contains(rendered.HTML, "nofollow") {
		t.Fatalf("link not hardened with rel attributes: %q", rendered.HTML)
	}
	if !strings.Contains(rendered.HTML, `target="_blank"`) {
		t.Fatalf("link missing target=_blank: %q", rendered.HTML)
	}
}

func TestRenderInboxMessageStripsStyles(t *testing.T) {
	rendered := render(t, `<p style="position:fixed;top:0">overlay</p><style>body{display:none}</style>`)
	if strings.Contains(rendered.HTML, "style") {
		t.Fatalf("style not stripped: %q", rendered.HTML)
	}
}
