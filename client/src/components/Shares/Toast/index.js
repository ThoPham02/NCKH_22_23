import { useEffect, useState } from "react";

import "./style.css";
import Error from "./ErrorToast";
import Success from "./SuccessToast";
import { useDispatch, useSelector } from "react-redux";
import { resultSelector } from "../../../store/selectors";
import Warning from "./WarningToast";

const toast = {
  0: {
    title: "Thành công!",
    message: "Đăng ký đề tài thành công!",
  },
  201: {
    title: "Không thành công!",
    message: "Bạn đã đăng ký đề tài khác!. Vui lòng kiểm tra lại!",
  },
  111: {
    title: "Không thành công!",
    message: "Đã xảy ra lỗi! Vui lòng kiểm tra lại!",
  },
};

const Toast = () => {
  const [showSuccess, setShowSuccess] = useState(false);
  const [showError, setShowError] = useState(false);
  const [showWarning, setShowWarning] = useState(false);
  const dispatch = useDispatch()
  let result = useSelector(resultSelector);

  const handleClose = () => {
    setShowSuccess(false);
    setShowError(false);
    setShowWarning(false);
  };

  console.log(result)

  useEffect(() => {
    if (result.code || result.code === 0) {
      switch (result.code) {
        case 0:
          setShowSuccess(true);
          break;
        case 201:
          setShowWarning(true);
          break;
        default:
          setShowWarning(true);
          break;
      }
    }
    // eslint-disable-next-line
  }, [dispatch, result]);

  return (
    <>
      <Success
        show={showSuccess}
        handleClose={handleClose}
        title={toast[0].title}
        message={toast[0].message}
        className="toasts"
      />
      <Error
        show={showError}
        handleClose={handleClose}
        title={toast[111].title}
        message={toast[111].message}
        className="toasts"
      />
      <Warning
        show={showWarning}
        handleClose={handleClose}
        title={toast[201].title}
        message={toast[201].message}
        className="toasts"
      />
    </>
  );
};

export default Toast;
