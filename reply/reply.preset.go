package reply

import (
	"fmt"
)

// AddPreset register reply value preset to client
//
// Example:
//
//	Client.AddPreset("RESOURCE_NOT_FOUND", func(rp *Reply, args ...any) *Reply {
//		resource := "resource"
//		if len(args) > 0 {
//			resourceArg, ok := args[0].(string)
//			if ok {
//				resource = resourceArg
//			}
//		}
//
//		return rp.Error("NOT_FOUND", resource+" not found.")
//	})
func (r *Client) AddPreset(name string, preset Preset) {
	r.presets[name] = preset
}

// AddSenderPreset register reply sender preset to client
//
// Example:
//
//	Client.AddSenderPreset("RESOURCE_NOT_FOUND", func(rp *Reply, args ...any) error {
//		resource := "resource"
//		if len(args) > 0 {
//			resourceArg, ok := args[0].(string)
//			if ok {
//				resource = resourceArg
//			}
//		}
//
//		return rp.Error("NOT_FOUND", resource+" not found.").FailJSON()
//	})
func (r *Client) AddSenderPreset(name string, preset SendPreset) {
	r.sendPresets[name] = preset
}

// UsePreset get preset from registered preset in client. return self instance and false if named preset don't exists
//
// Example:
//
//	rp, exists := rp.UsePreset("RESOURCE_NOT_FOUND", "user")
//	if exists {
//		rp.Info("user with id " + userId + " not found").FailJSON()
//	} else {
//		rp.Error("NOT_FOUND", "user not found").
//			Info("user with id " + userId + " not found").
//			Debug("preset RESOURCE_NOT_FOUND not found").
//			FailJSON()
//	}
func (r *Reply) UsePreset(name string, args ...any) (instance *Reply, exists bool) {
	preset, exists := r.c.presets[name]
	if !exists {
		return r, false
	}

	return preset(r, args...), true
}

// SendPreset sends response with registered sender preset. return joined [ErrPresetNotFound] if preset not exist
//
// Example:
//
//	err := rp.SendPreset("RESOURCE_NOT_FOUND", "user")
//	if errors.Is(reply.ErrPresetNotFound) {
//		rp.Error("NOT_FOUND", "user not found.").
//			Debug("preset RESOURCE_NOT_FOUND not found").
//			FailJSON()
//	} else {
//		return err
//	}
func (r *Reply) SendPreset(name string, args ...any) (err error) {
	preset, exists := r.c.sendPresets[name]
	if !exists {
		// join to compare with errors.Is
		return fmt.Errorf("%w, name: %s", ErrPresetNotFound, name)
	}

	return preset(r, args...)
}
