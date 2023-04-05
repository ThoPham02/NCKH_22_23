import Pagination from 'react-bootstrap/Pagination';
function Pagination1({ setPagi, currentPage, numberItem }) {
  let items = [];
  for (const number = 1; number <= 5; number++) {
    items.push(
      <Pagination.Item key={number} active={number === active}>
        {number}
      </Pagination.Item>,
    );
  }
  return (
    <Pagination size="sm">{items}</Pagination>
  )
}
export default Pagination1



