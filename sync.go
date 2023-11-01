package main

func setActiveLine(elapsed float64, lrcLines *[]Lyric) {
	list := *lrcLines

	if len(list) == 1 {
		return
	}

	for i := range list {
		line := &list[i]
		nextLine := line

		isLast := i+1 >= len(list)

		if !isLast {
			nextLine = &list[i+1]
		}

		line.Active = (elapsed >= line.Timestamp && elapsed < nextLine.Timestamp) || isLast && elapsed >= list[len(list)-1].Timestamp
	}
}
