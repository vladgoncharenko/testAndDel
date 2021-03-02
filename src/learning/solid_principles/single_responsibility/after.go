package single_responsibility

import "image"

// NavigateTo applies any required changes to the drone's speed vector
// so that its eventual position matches dst.
func (d *Drone) NavigateToR(dst Vec3) error { return nil }

// Position returns the current drone position vector.
func (d *Drone) PositionR() Vec3 { return Vec3{} }

// Position returns the current drone speed vector.
func (d *Drone) SpeedR() Vec3 { return Vec3{} }

// CaptureImage records and returns an image of the drone's field of
// view using the on-board drone camera.
func (d *Drone) CaptureImage() (*image.RGBA, error) { return nil, nil }

//In a separate file (possibly in a different package too), we would define the MobileNet type,
//which contains the implementation for our target detector:

// MobileNet performs target detection for drones using the
// SSD MobileNet V1 NN.
// For more info on this model see:
// https://github.com/tensorflow/models/tree/master/research/object_detection
//type MobileNet {
// various attributes...
//}

// DetectTargets captures an image of the drone's field of view and feeds
// it to a neural network to detect and classify interesting nearby
// targets.
//func (mn *MobileNet) DetectTargets(d *drone.Drone) ([]*Target, error){
//...
//}
