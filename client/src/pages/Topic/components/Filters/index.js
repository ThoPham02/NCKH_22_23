import { useState } from "react"
import { BsSearch } from "react-icons/bs"
import "./search__input.css"
import { useDispatch } from "react-redux";
import filtersSlice from "./filtersSlice"

function SelectInput() {
    const dispatch = useDispatch();
    const [searchText, setSearchText] = useState('');

    const handleSearchTextChange = (e) => {
        setSearchText(e.target.value);
        dispatch(
            filtersSlice.actions.searchFilterChange(e.target.value)
        )
    }
    return (

        <div id="icon--search">
            <input placeholder="Tìm kiếm nhanh..."
                value={searchText}
                onChange={handleSearchTextChange}
            />
            <button>
                <BsSearch></BsSearch>
            </button>
        </div>
    );
}
export default SelectInput;
