package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/AlickZ/fyne_test/UI"
)

func main() {
	// -----------------------------------------新建APP 实例----------------------------------
	a := app.New()
	// ------------------------------------------新建窗口-------------------------------------
	w := a.NewWindow("Linux 路由检查")
	// ---------------------------------禁止用户调整窗口大小 并监听界面变化-------------------------
	w.Resize(fyne.NewSize(1000, 600))
	w.SetFixedSize(true) // 禁止用户调整窗口大小
	// 加载自定义图标
	icon, _ := fyne.LoadResourceFromPath("static/images/bitbug_favicon.ico")
	w.SetIcon(icon)
	// -----------------------------------设置窗口屏幕居中--------------------------------------
	w.CenterOnScreen()
	// ------------------------------------设置内容 加载其他组件---------------------------------
	ap := UI.ManinUI{}
	w.SetContent(ap.OutINFO())

	// 获取窗口大小
	// -----------------------------系统托盘---------------------------------------------------
	if desk, ok := a.(desktop.App); ok {

		// 新建展示按钮
		h := fyne.NewMenuItem("打开主页面", func() {
			w.Show()
		})
		h.Icon = icon // 按钮单独设置图标
		m := fyne.NewMenu("MyApp", h)
		// 按钮添加到系统托盘
		desk.SetSystemTrayMenu(m)
		// 设置系统托盘图标
		desk.SetSystemTrayIcon(icon)
	}

	// 设置关闭 方法为隐藏
	w.SetCloseIntercept(func() {
		w.Hide()
	})
	// -------------------------------------运行程序--------------------------------------------
	w.ShowAndRun()
}
