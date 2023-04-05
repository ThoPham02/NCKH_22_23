import Pagination from 'react-bootstrap/Pagination';
function Pagination1({setPagi, currentPage, numberItem}) {
  let content = <></>
 
  if (currentPage < 5) {
    for ( let i = 1; i<5; i++) {
      document.getElementById("Pagination").innerHTML = `<li>${i}<li>`
    }
  }

  


    return(
        <Pagination >
        <Pagination.Item onClick={(e) => {setPagi(1)}} id='Pagination'>{1}</Pagination.Item>
          {content}
      </Pagination>
    )
}
export default Pagination1