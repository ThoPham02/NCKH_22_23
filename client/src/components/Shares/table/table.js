import Table from 'react-bootstrap/Table'

function BasicExample() {
  return (
    <Table striped bordered hover>
      <thead>
        <tr>
          <th>Số thứ tự</th>
          <th>Thông tin đề tài</th>
          <th>thông tin chi tiết</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>1</td>
          <td>Mark</td>
          <td>Otto</td>
          
        </tr>
        <tr>
          <td>2</td>
          <td>Jacob</td>
          <td>Thornton</td>
          
        </tr>
        <tr>
          <td>3</td>
          <td>Larry the Bird</td>
          <td>@twitter</td>
        </tr>
      </tbody>
    </Table>
  );
}

export default BasicExample;