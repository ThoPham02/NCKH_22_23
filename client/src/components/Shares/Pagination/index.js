import Pagination from "react-bootstrap/Pagination";

import "./pagination.css";

const PaginationCustom = ({ setPagi, currentPage, total, limit }) => {
  let num = Math.ceil(total / limit);
  let pagiList = [];
  for (let i = 0; i < num; i++) {
    pagiList.push(i + 1);
  }
  return (
    <div className="pagination">
      
      <Pagination size="sm">
      <Pagination.Prev onClick={()=>setPagi(currentPage-1)} disabled={currentPage===1}/>
        {pagiList.map((number) => {
          return (
            <Pagination.Item
              key={number}
              active={number === currentPage}       
              onClick={(e) => setPagi(e.target.innerHTML*1)} 
            >
              {number}
            </Pagination.Item>
          );
        })}
      <Pagination.Next onClick={()=>setPagi(currentPage+1)} disabled={currentPage===num}/>
      </Pagination>

    </div>
  );
};

export default PaginationCustom;
