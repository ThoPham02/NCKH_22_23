import { Table } from "react-bootstrap";
import { useDispatch } from "react-redux";
import { useRef } from "react";

import "./style.css"
import Card from "../../../components/Shares/Card";
import { TopicSearch } from "../../../components/Shares/Search";

const Topic = () => {
    useDispatch();
    const searchRef = useRef("");
    const statusRef = useRef(0);
    const dateFromRef = useRef("");
    const dateToRef = useRef("");
    const departmentRef = useRef(0);

    const handleSubmitForm = () => {
        
    }
    return (
        <>
            <Card title={"Danh sách đề tài"}>
                <TopicSearch
                    handleSubmitForm={handleSubmitForm}
                    searchRef={searchRef}
                    statusRef={statusRef}
                    dateFromRef={dateFromRef}
                    dateToRef={dateToRef}
                    departmentRef={departmentRef}
                />
                <Table bordered hover size="sm" className="topic-table">

                </Table>
            </Card>
        </>
    )
};

export default Topic;
