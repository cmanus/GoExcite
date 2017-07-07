// Objective-C API for talking to github.com/cmanus/GoExcite/Golang Go package.
//   gobind -lang=objc github.com/cmanus/GoExcite/Golang
//
// File is generated by gobind. Do not edit.

#ifndef __GoExcite_H__
#define __GoExcite_H__

@import UIKit;
@import Foundation;
#include "Universe.objc.h"


@class GoExciteMainViewController;

@interface GoExciteMainViewController : UIViewController {
}
@property(strong, readonly) id _ref;

- (id)initWithRef:(id)ref;
- (id)init;
- (UIViewController *)uiViewController;
- (void)setUIViewController:(UIViewController *)v;
- (void)didReceiveMemoryWarning;
- (void)viewDidLoad;
@end

#endif