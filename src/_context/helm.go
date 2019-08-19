package _context

var HelmApp helmApp

func InitApplicationContext(service, image, tag string) *helmApp {

	HelmApp = helmApp{
		Service: service,
		Image: image,
		Tag: tag,
	}

	return &HelmApp
}

type helmApp struct {
	Service string
	Image   string
	Tag     string
}