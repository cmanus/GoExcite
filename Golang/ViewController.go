// +build ios
package goExcite

import (
	//"fmt"
	"ObjC/UIKit"
	//"ObjC/Foundation"
	gopkg "ObjC/GoExcite"
	
)

type MainViewController struct {
	UIKit.UIViewController
}

func (vc *MainViewController) ViewDidLoad(self gopkg.MainViewController) {
	self.Super().ViewDidLoad()
	//fmt.Println(self.View().Bounds())
	//webview := UIKit.UIWebView.NewWithFrame(self.View().Bounds())
	//url := Foundation.NSURL.URLWithString("https://google.com")
	//request := Foundation.NSURLRequest.RequestWithURL(url)
	//webview.LoadRequest(request)
	//self.View().AddSubView(webview)
	
}
func (vc *MainViewController) DidReceiveMemoryWarning(self gopkg.MainViewController) {
	self.Super().DidReceiveMemoryWarning()
}

