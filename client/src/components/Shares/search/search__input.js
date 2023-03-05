import {BsSearch} from "react-icons/bs"

function SelectInput(){
    return(
        <div id="icon--search">
          <input placeholder="Tìm kiếm nhanh..."/>
            <button>
              <BsSearch></BsSearch>
            </button>
        </div>
    );
}
export default SelectInput;