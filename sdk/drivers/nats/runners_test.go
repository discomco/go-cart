package nats

import (
	"context"
	"fmt"
	"log"
)

func publishWorker(ctx context.Context, ins chan []byte) func() error {
	return func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				var bus INATSBus
				var err error
				i := 0
				err = testEnv.Invoke(func(newBus BusFtor) {
					bus, err = newBus()
				})
				if err != nil {
					log.Fatal(err)
					return err
				}
				for {
					in := <-ins
					i = i + 1
					testLogger.Infof("Publishing message %d", i)
					bus.Publish(ctx, TEST_TOPIC, []byte(fmt.Sprintf(`{"message": "[{%+v}] #{%+v}"}`, string(in), i)))
				}
			}
		}
	}
}

func listenWorker(ctx context.Context, outs chan []byte) func() error {
	return func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			var bus INATSBus
			var err error
			err = testEnv.Invoke(func(newBus BusFtor) {
				bus, err = newBus()
			})
			if err != nil {
				log.Fatal(err)
				return err
			}
			bus.Listen(ctx, TEST_TOPIC, outs)
		}
		return nil
	}
}

//func respondWorker(ctx context.Context, outs chan []byte) func() error {
//	return func() error {
//		select {
//		case <-ctx.Done():
//			return ctx.Err()
//		default:
//			var bus INATSBus
//			var err error
//			testEnv.Invoke(func(newBus BusFtor) {
//				bus, err = newBus()
//			})
//			if err != nil {
//				log.Fatal(err)
//				return err
//			}
//			bus.Respond(ctx, TEST_TOPIC, outs)
//		}
//		return nil
//	}
//}
