package common

var MapDegree = map[int64]string{
	1:  "SV",
	2:  "ThS",
	4:  "TS",
	8:  "PGS",
	16: "GS",
}

var MapStatus = map[int64]string{
	1:    "Hủy",
	2:    "Đề Xuất",
	4:    "Đăng Ký",
	8:    "Chờ Phê Duyệt",
	16:   "Đang Thực Hiện",
	32:   "Báo Cáo Tiến Độ",
	64:   "Báo Cáo Tiểu Ban",
	128:  "Báo Cáo Cấp Trường",
	256:  "Hoàn Thành",
	512:  "Hoàn thành Quá Hạn",
	1024: "Không Hoàn Thành",
}
