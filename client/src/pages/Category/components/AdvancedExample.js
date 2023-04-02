import { useState } from 'react';
import categoryList from '../categoryList';
import "./AdvancedExample.css"
import Pagination1 from '../../../components/Shares/Pagination';
function AdvancedExample() {
  var data = categoryList;
  const [currentPage, setCurrentPage] = useState(1);
  const perPage = 3; 
  const startIndex = (currentPage - 1) * perPage;
  const endIndex = startIndex + perPage;
  const currentItems = data.slice(startIndex, endIndex);  
  return (
    <>
      {
      currentItems.map((todo, index) => {
        return (
          <div key={index} >
            <p>{index + 1}</p>
            <p>{todo.nameCategory}</p>
            <a href={todo.url}>url</a>
          </div>
        )
      })
    }
   <Pagination1 setPagi={setCurrentPage} currentPage={currentPage} numberItem={10}></Pagination1>
    </>
  );
}

export default AdvancedExample;