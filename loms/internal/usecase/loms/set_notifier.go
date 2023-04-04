package loms

import "route256/loms/internal/notifier"

func (u *useCase) SetNotifier(notifier notifier.Notifier) {
	u.notifier = notifier
}
