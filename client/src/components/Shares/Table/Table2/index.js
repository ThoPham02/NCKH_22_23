import Table from 'react-bootstrap/Table'
import TableHeader from './TableHeader'
import "./style.css"

function TableTopic2({ listHead, listItem, listKey }) {
  return (
    <div>
      <Table striped bordered hover>
        <TableHeader listKey={listHead} />
        <tbody>
        {listItem.map((item, index) => (
          <tr key={index}>
            <td>{index + 1}</td>
            {listKey.map((key, index) => (
              <td key={index}>{item[key]}</td>
            ))}
            <td>
              <button style={{borderTop:"none",borderLeft:"none",borderRight:"none",width:"70px"}}/*onClick={() => handleRegister(item)}*/>Đăng kí</button>
            </td>
          </tr>
        ))}
      </tbody>
      </Table>
      <div className="prop">
        <p>đăng kí đề tài</p>
      </div>
    </div>
  );
}

export default TableTopic2;