import { createSlice } from "@reduxjs/toolkit";

const StageSlice = createSlice({
  name: "stage",
  initialState: [
    {
      id: 1,
      name: "Đề xuất",
      dateFrom: "21/05/2022",
      dateTo: "21/07/2022",
      description:
        "Giảng viên thực hiện đề xuất các đề tài. Các khoa rà soát, kiểm tra và đăng ký các đề tài để sinh viên đăng ký",
    },
    {
      id: 2,
      name: "Đăng ký",
      dateFrom: "21/07/2022",
      dateTo: "21/08/2022",
      description:
        "Sinh viên thực hiện đăng ký đề tài. Các khoa rà soát, kiểm tra và đăng ký các đề tài để sinh viên bắt dầu nghiên cứu",
    },
    {
      id: 3,
      name: "Thực hiện",
      dateFrom: "21/08/2022",
      dateTo: "21/5/2023",
      description:
        "Nhóm sinh viên thực hiện nghiên cứu đề tài đã đăng ký dưới sự hướng dẫn của giảng viên hướng dẫn",
    },
    {
      id: 4,
      name: "Nghiệm thu tại bộ môn",
      dateFrom: "21/05/2023",
      dateTo: "28/05/2023",
      description:
        "Các bộ môn chủ quản tiến hành đánh giá, nghiệm thu đề tài nghiên cứu khoa học sinh viên(NCKH SV). Văn phòng khoa tập hợp các đề tài sau khi nghiệm thu và nộp về phòng Khoa học Công nghệ(KHCN) trước 17h00 ngày 28/05/2023.",
    },
    {
      id: 5,
      name: "Nghiệm thu tại tiểu ban",
      dateFrom: "",
      dateTo: "",
      description:
      "Coming soon...",
    },
    {
      id: 6,
      name: "Báo cáo cấp trường",
      dateFrom: "",
      dateTo: "",
      description: "Coming soon...",
    },
    {
      id: 7,
      name: "Kết thúc",
      dateFrom: "",
      dateTo: "",
      description: "Coming soon...",
    },
  ],
  reducers: {
    updateStage: (state, action) => {
      state[action.payload.stageDetail] = {
        id: state[action.payload.stageDetail].id,
        name: state[action.payload.stageDetail].name,
        description: action.payload.text,
        dateFrom: action.payload.dateFrom,
        dateTo: action.payload.dateTo,
      };
    },
  },
});

export default StageSlice;
export const StageReducer = StageSlice.reducer;
export const StageAction = StageSlice.actions
