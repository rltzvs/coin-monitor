package telegram

type UserUsecaseImpl struct{}

func (u *UserUsecaseImpl) StartAutoUpdate(userId int, interval int) error {
	// Заглушка
	return nil
}

func (u *UserUsecaseImpl) StopAutoUpdate(userId int) error {
	// Заглушка
	return nil
}
