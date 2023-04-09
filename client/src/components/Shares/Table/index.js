import Table from 'react-bootstrap/Table'

import TableHeader from './TableHeader'
import TableBody from './TableBody';

function TableTopic({listHead, listItem, listKey}) {
  return (
    <Table striped bordered hover>
      <TableHeader listKey={listHead}/>
      <TableBody listItem={listItem} listKey={listKey}/>
    </Table>
  );
}

export default TableTopic;