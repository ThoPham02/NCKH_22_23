import Pagination from "react-bootstrap/Pagination";

import "./pagination.css";

const Pagination1 = ({ setPagi, currentPage, total, limit }) => {
  let num = Math.ceil(total / limit);
  let pagiList = [];
  for (let i = 0; i < num; i++) {
    pagiList.push(i + 1);
  }
  return (
    <div className="pagination">
      <Pagination>
        {pagiList.map((number) => {
          return (
            <Pagination.Item
              key={number}
              active={number === currentPage}
              onClick={(e) => setPagi(e.target.value)}
            >
              {number}
            </Pagination.Item>
          );
        })}
      </Pagination>
    </div>
  );
};

export default Pagination1;
