import Form from 'react-bootstrap/Form';
import "./search__form.css"


function SelectBasicExample() {
  function myFunction() {
    document.getElementById('topic__search--table').style.display = "block"
  }
  return (
    <>
    <button onClick={myFunction}>
      lọc chi tiết
    </button>
    <div id="topic__search--table">
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
      <label htmlFor="start">Từ ngày:</label>
        <input type="date" id="start" name="trip-start"/>
      <label htmlFor="start">Đến ngày:</label>
        <input type="date" id="start" name="trip-start"/>
    </div>

    
    </>
    
  );
}





export default SelectBasicExample;