import Form from 'react-bootstrap/Form';
import { useState } from 'react';
import "./search__form.css"


function SelectBasicExample(props) {
  var [open, setOpen] = useState(props.openSearch)
  function myFunction() {
    setOpen(preState => open = !preState)
  }

  return (
    <>

      <button onClick={myFunction}>
        lọc chi tiết
      </button>
      <div id="topic__search--table" style={open?{display:"block"}:{display:"none"}}>
      <Form.Select aria-label="Default select example">
        <option>Open this select menu</option>
        <option value="1">One</option>
        <option value="2">Two</option>
        <option value="3">Three</option>
      </Form.Select>
      <Form.Select aria-label="Default select example">
        <option>Open this select menu</option>
        <option value="1">One</option>
        <option value="2">Two</option>
        <option value="3">Three</option>
      </Form.Select>
      <Form.Select aria-label="Default select example">
        <option>Open this select menu</option>
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