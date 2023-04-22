// import Col from "react-bootstrap/Col";
// import Form from "react-bootstrap/Form";
// import Row from "react-bootstrap/Row";
// import Button from 'react-bootstrap/Button';
// import { useState } from "react";
// import { useDispatch } from "react-redux"

import DateFrom from "./DateFrom";
import DateTo from "./DateTo";
import Department from "./Department";
import Faculty from "./Faculty";
import Status from "./Status";
import Word from "./Word";

// import "./search.css";
// import SearchWord from "./SearchWord/indes";
// import SearchStatus from "./SearchStatus";
// import SearchDepartment from "./SearchDepartment";
// import SearchFaculity from "./SearchFaculity";
// import SearchDate from "./SearchDate";
// import FilterSlice from "./SearchSlice";

// export const Search = () => {
//   const [searchWord, setSearchWord] = useState("")
//   const [dateFrom, setDateFrom] = useState("")
//   const [dateTo, setDateTo] = useState("")
//   const [searchStatus, setSearchStatus] = useState(0)
//   const [searchDepartment, setSearchDepartment] = useState(0)
//   const [searchFaculity, setSearchFaculity] = useState(0)

//   const dispatch = useDispatch()

//   const handleSearch = (e) => {
//     e.preventDefault()
    
//     dispatch(FilterSlice.actions.searchFilterChange({
//       search: searchWord,
//       dateFrom: dateFrom,
//       dateTo: dateTo,
//       status: searchStatus * 1,
//       department: searchDepartment * 1,
//       faculity: searchFaculity * 1
//   }))
//   }

//   return (
//     <Form className="search-form" onSubmit={handleSearch}>
//       <Row className="search-row">
//         <Col xs={8}>
//           <SearchWord setSearchWord={setSearchWord}/>
//         </Col>
//         <Col>
//           <SearchStatus setSearchStatus={setSearchStatus}/>
//         </Col>
//       </Row>
//       <Row className="search-row">
//         <Col>
//           <SearchDepartment setSearchDepartment={setSearchDepartment}/>
//         </Col>
//         <Col>
//           <SearchFaculity setSearchFaculity={setSearchFaculity}/>
//         </Col>
//         <Col>
//           <SearchDate setDateFrom={setDateFrom} setDateTo={setDateTo}/>
//         </Col>
//       </Row>

//       <Button variant="primary" type="submit">
//         Tìm kiếm
//       </Button>
//     </Form>
//   );
// };

export const SearchWord = Word;
export const SearchStatus = Status;
export const SearchDateTo = DateTo;
export const SearchDateFrom = DateFrom;
export const SearchFaculty = Faculty;
export const SearchDepartment = Department;