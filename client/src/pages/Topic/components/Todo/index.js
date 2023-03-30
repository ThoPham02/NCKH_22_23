import Table from 'react-bootstrap/Table'
function TodoTable({ todoList }) {
  return (
    <Table striped bordered hover>
      <thead>
        <tr>
          <th>Số thứ tự</th>
          <th>Thông tin đề tài</th>
          <th>thông tin chi tiết</th>
        </tr>
      </thead>
      {todoList.map((todo, index) => {
        return (
          <tbody key={index}>
            <tr>
              <td>{index+1}</td>
              <td>
                <span>{todo.name}</span>
                <br></br>
                <span>{todo.completed}</span>
              </td>
              <td></td>
            </tr>
          </tbody>
        )
      })}
    </Table>
  )
}

export default TodoTable;