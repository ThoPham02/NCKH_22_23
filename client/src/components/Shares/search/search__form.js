import Form from 'react-bootstrap/Form';
import { useState } from 'react';
import "./search__form.css"
<>
</>


function SelectBasicExample(props) {
  var [open, setOpen] = useState(props.openSearch)
  function myFunction() {
    setOpen(preState => open = !preState)
  }

  return (
    <>
      <button onClick={myFunction} id= "filter">
        lọc chi tiết
      </button>
      <div id="topic__search--table" style={open?{display:"block"}:{display:"none"}}>
      <Form.Select aria-label="Default select example">
        <option value="tất cả trạng thái">tất cả trạng thái</option>
        <option value="hoàn thành">hoàn thành</option>
        <option value="quá hạn">quá hạn</option>
        <option value="đang thực hiện">đang thực hiện</option>
      </Form.Select>
      <Form.Select aria-label="Default select example">
        <option value="tất cả bộ môn">tất cả bộ môn</option>
        <option value="1">One</option>
        <option value="2">Two</option>
        <option value="3">Three</option>
      </Form.Select>
      <Form.Select aria-label="Default select example">
        <option value = "tất cả khoa">tất cả khoa</option>
        <option value="1">One</option>
        <option value="2">Two</option>
        <option value="3">Three</option>
      </Form.Select>
      <div id="selectDate">
        <label htmlFor="start">Từ ngày:</label>
        <input type="date" id="start" name="trip-start" />
        <label htmlFor="start">Đến ngày:</label>
        <input type="date" id="start" name="trip-start" />
      </div>
    </div>

    
    </>
    
  );
}

export default SelectBasicExample;