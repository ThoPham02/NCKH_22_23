import Pagination from "react-bootstrap/Pagination";

import "./pagination.css";

const PaginationCustom = ({ setPagi, pagi, total, limit, style }) => {
  let num = Math.ceil(total / limit);
  let pagiList = [];
  for (let i = 0; i < num; i++) {
    pagiList.push(i + 1);
  }
  return (
    <div className="pagination" style={style}>
      <Pagination size="sm">
      <Pagination.Prev onClick={()=>setPagi(pagi-1)} disabled={pagi===1}/>
        {pagiList.map((number) => {
          return (
            <Pagination.Item
              key={number}
              active={number === pagi}       
              onClick={(e) => setPagi(e.target.innerHTML*1)} 
            >
              {number}
            </Pagination.Item>
          );
        })}
      <Pagination.Next onClick={()=>setPagi(pagi+1)} disabled={pagi===num}/>
      </Pagination>

    </div>
  );
};

export default PaginationCustom;
