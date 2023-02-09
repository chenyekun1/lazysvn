package tui

import (
	"container/list"
	"fmt"
	"log"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func passSearch(searchbar *tview.TextView, main *tview.Table, screen *TuiScreen, LOG *log.Logger) {
	for item := screen.searchRes.Front(); item != nil; item = item.Next() {
		info := item.Value.(searchInfo)
		LOG.Println(fmt.Sprintf("erase %d", info.row))
		main.GetCell(info.row, 0).SetBackgroundColor(tcell.ColorBlack)
	}
	screen.searchRes = list.New()
	for i := 0; i <main.GetRowCount(); i++ {
		fileName := main.GetCell(i, 0).Text
		if strings.Contains(fileName, searchbar.GetText(true)[1:]) {
			main.GetCell(i, 0).SetBackgroundColor(tcell.ColorYellow)
			screen.searchRes.PushBack(searchInfo{row:i})
		}
	}
	for item := screen.searchRes.Front(); item != nil; item = item.Next() {
		info := item.Value.(searchInfo)
		LOG.Println(fmt.Sprintf("erase2 %d", info.row))
	}
}

func passSearchJump(main *tview.Table, screen *TuiScreen, LOG *log.Logger) {
	idx, _ := main.GetSelection()
	if screen.searchRes.Len() <= 0 {
		return
	}
	searchRes := screen.searchRes
	for item := searchRes.Front(); item != nil; item = item.Next() {
		info := item.Value.(searchInfo)
		if idx < info.row {
			main.Select(info.row, 0)
			return
		}
	}
	first := searchRes.Front().Value.(searchInfo)
	main.Select(first.row, 0)
}
