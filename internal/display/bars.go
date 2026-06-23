package display

func PriorityBars(p int) string {
	switch p {
	case 0:
		return "▁▄█"
	case 1:
		return "▁▄"
	case 2:
		return "▁"
	default:
		return ""
	}
}
