import Table from 'react-bootstrap/Table'
import TableHeader from './TableHeader'
import TableBody from './TableBody';

function TableTopic2({ listHead, listItem, listKey }) {
  return (
    <div>
      <Table striped bordered hover>
        <TableHeader listKey={listHead} />
        <TableBody listItem={listItem} listKey={listKey} />
      </Table>
    </div>
  );
}

export default TableTopic2;