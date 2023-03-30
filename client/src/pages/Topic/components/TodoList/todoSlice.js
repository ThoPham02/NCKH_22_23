// const initState = [
//     { id: 1, name: 'NCKH CNTT', completed: 'hoàn thành' },
//     { id: 2, name: 'NCKH CNTT 2', completed: 'quá hạn' },
//     { id: 3, name: 'NCKH CNTT 3', completed: 'đang thực hiện' }
// ]

// const todoListReducer = (state = initState, action) => {
//     switch (action.type) {
//         case 'todoList/addTodo':
//             return [
//                 ...state.todoList,
//                 action.payload
//             ]

//         default:
//             return state;
//     }
// }

// export default todoListReducer;

import { createSlice } from "@reduxjs/toolkit";

export default createSlice({
    name: 'todoList',
    initialState:
        [
            { id: 1, name: 'NCKH CNTT', completed: 'hoàn thành' },
            { id: 2, name: 'NCKH CNTT 2', completed: 'quá hạn' },
            { id: 3, name: 'NCKH CNTT 3', completed: 'đang thực hiện' }

        ],
    reducers: {
        addTodo: (state, action) => {
            state.push(action.payload);
        },
    }

})