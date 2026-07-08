package repository

import (
	"context"

	"ivpn.net/email/api/internal/model"
)

func (d *Database) GetPlan(ctx context.Context, id string) (model.Plan, error) {
	var plan model.Plan
	err := d.Client.First(&plan, "id = ?", id).Error
	return plan, err
}

func (d *Database) GetPlanByName(ctx context.Context, name string) (model.Plan, error) {
	var plan model.Plan
	err := d.Client.First(&plan, "name = ?", name).Error
	return plan, err
}

func (d *Database) GetActivePlans(ctx context.Context) ([]model.Plan, error) {
	var plans []model.Plan
	err := d.Client.Where("is_active = ?", true).Order("sort_order asc").Find(&plans).Error
	return plans, err
}

func (d *Database) GetAllPlans(ctx context.Context) ([]model.Plan, error) {
	var plans []model.Plan
	err := d.Client.Order("sort_order asc").Find(&plans).Error
	return plans, err
}

func (d *Database) PostPlan(ctx context.Context, plan model.Plan) error {
	return d.Client.Create(&plan).Error
}

func (d *Database) UpdatePlan(ctx context.Context, plan model.Plan) error {
	return d.Client.Save(&plan).Error
}

func (d *Database) DeletePlan(ctx context.Context, id string) error {
	return d.Client.Model(&model.Plan{}).Where("id = ?", id).Update("is_active", false).Error
}
