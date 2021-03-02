package single_responsibility

type Drone struct {
}

type Vec3 struct {
}

type Target struct {
}

// NavigateTo applies any required changes to the drone's speed
// vector so that its eventual position matches dst.
func (d *Drone) NavigateTo(dst Vec3) error {
	return nil
}

// Position returns the current drone position vector.
func (d *Drone) Position() Vec3 {
	return Vec3{}
}

// Position returns the current drone speed vector.
func (d *Drone) Speed() Vec3 { return Vec3{} }

// DetectTargets captures an image of the drone's field of view (FoV) using
// the on-board camera and feeds it to a pre-trained SSD MobileNet V1 neural
// network to detect and classify interesting nearby targets. For more info
// on this model see:
// https://github.com/tensorflow/models/tree/master/research/object_detection
func (d *Drone) DetectTargets() ([]*Target, error) {
	var res []*Target
	return res, nil
}
