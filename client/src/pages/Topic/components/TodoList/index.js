import "./style.css"
import { CgAdd } from "react-icons/cg"
import Form from 'react-bootstrap/Form';
import { useDispatch, } from 'react-redux';
import { v4 as uuidv4 } from 'uuid';
import { useState } from "react";

import todoListSlice from "./todoSlice"


function AddTodo() {

    const [todoName, setTodoName] = useState('');
    const [completed, setCompleted] = useState('hoàn thành');



    const dispatch = useDispatch();
    const handleAddButtonClick = () => {
        dispatch(
            todoListSlice.actions.addTodo({
            id: uuidv4(),
            name: todoName,
            completed: completed
        }));

        setTodoName('');
        setCompleted('hoàn thành');
    }

    const handleCompletedChange = (e) => {
        setCompleted(e.target.value)
    }
    const handleInputChange = (e) => {
        setTodoName(e.target.value);
    }
    return (
        <div id="addTodo">
            <input placeholder="thêm đề tài..." value={todoName} onChange={handleInputChange} />
            <Form.Select aria-label="Default select example" size="sm" id="state"
                onChange={handleCompletedChange}>
                <option value='hoàn thành'>hoàn thành</option>
                <option value='chưa hoàn thành'>chưa hoàn thành</option>
                <option value='đang hoàn thành'>đang hoàn thành</option>
            </Form.Select>
            <button onClick={handleAddButtonClick}>
                <CgAdd></CgAdd>
            </button>
        </div>
    );
}
export default AddTodo;