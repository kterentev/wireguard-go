package device

import (
	"time"
)

func (device *Device) RoutineTimeUpdater() {
	defer func() {
		device.log.Verbosef("Routine: Time updater - stopped")
		device.state.stopping.Done()
	}()

	device.log.Verbosef("Routine: Time updater - started")
	device.time.now = time.Now()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			device.time.now = t
		case <-device.time.c:
			return
		}
	}
}
