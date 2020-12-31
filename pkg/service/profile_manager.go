package service

import "testGinandGorm/pkg/service/profile"

type ProfileRuntime interface {
	Push(profile.CreateContext) error
	QueryById(profile.QueryContext) error
	QueryByName(profile.QueryByNameContext) error
	QueryMutis(profile.QueryObjectsContext) error
	UpdateById(profile.UpdateContext) error
	UpdateByNo(profile.UpdateContext) error
	Delete(profile.DeleteContext) error
}

type ProfileManager struct {
	profileRuntime ProfileRuntime
}

func NewProfileManager(profileRuntime ProfileRuntime) *ProfileManager {
	return &ProfileManager{profileRuntime: profileRuntime}
}

func (m *ProfileManager) PushProfile(ctx profile.CreateContext) error {
	return m.profileRuntime.Push(ctx)
}

func (m *ProfileManager) QueryProfileById(ctx profile.QueryContext) error {
	return m.profileRuntime.QueryById(ctx)
}

func (m *ProfileManager) QueryProfileByName(ctx profile.QueryByNameContext) error {
	return m.profileRuntime.QueryByName(ctx)
}

func (m *ProfileManager) QueryProfiles(ctx profile.QueryObjectsContext) error {
	return m.profileRuntime.QueryMutis(ctx)
}

func (m *ProfileManager) UpdateProfileById(ctx profile.UpdateContext) error {
	return m.profileRuntime.UpdateById(ctx)
}

func (m *ProfileManager) UpdateProfileByNo(ctx profile.UpdateContext) error {
	return m.profileRuntime.UpdateByNo(ctx)
}

func (m *ProfileManager) DeleteProfile(ctx profile.DeleteContext) error {
	return m.profileRuntime.Delete(ctx)
}
