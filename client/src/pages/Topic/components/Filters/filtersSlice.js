// const initState = {
//     search: '',
//     status: 'All',
// }

// const filtersReducer = (state = initState, action) => {
//     switch (action.type) {
//         case 'filters/searchFilterChange':
//             return {
//                 ...state,
//                 search: action.payload
//             }
//         default:
//             return state;
//     }
// }

// export default filtersReducer;

import { createSlice } from "@reduxjs/toolkit";

export default createSlice({
    name: 'filters',
    initialState: {
        search: '',
        status: 'All',
    },
    reducers: {
        searchFilterChange(state, action) {
            state.search = action.payload;
        }
        //
        //
    }
})
/*
khi khai báo reducers ở dòng số 28 nó sẽ trả về
1 action creator
(payload) => {
return {
    type: 'name/tên reducer' tức là
    type: 'filters/searchFilterChange'
    payload: payload
}
}
*/