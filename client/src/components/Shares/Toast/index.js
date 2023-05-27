import { useEffect, useState } from 'react';
import { Toast } from 'react-bootstrap';

function ToastConfirm() {
  const [showA, setShowA] = useState(false);

  useEffect(() => {
    if (showA) {
      setTimeout(() => {
        setShowA(false);
      }, 3000);
    }
  }, [showA]);

  const toggleShowA = () => setShowA(!showA);

  return (
    <div className="toast-container">
      <Toast
        show={showA}
        onClose={toggleShowA}
        className={showA ? 'toast-slide-in' : 'toast-slide-out'}
      >
        <Toast.Header>
            <strong className="me-auto">Thao tác thành công</strong>
          </Toast.Header>
      </Toast>
    </div>
  );
}

export default ToastConfirm;
