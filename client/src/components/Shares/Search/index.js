import { useState } from "react";
import { Button, Form } from "react-bootstrap";

import DateFrom from "./DateFrom";
import DateTo from "./DateTo";
import Department from "./Department";
import Faculty from "./Faculty";
import Status from "./Status";
import Word from "./Word";
import Loading from "../Loading";
import Event from "./Event";

export const SearchWord = Word;
export const SearchStatus = Status;
export const SearchDateTo = DateTo;
export const SearchDateFrom = DateFrom;
export const SearchFaculty = Faculty;
export const SearchDepartment = Department;

export const TopicSearch = (props) => {
    const { handleSubmitForm, searchRef, departmentRef, statusRef, dateFromRef, dateToRef, isLoading } = props

    const [faculty, setFaculty] = useState(0)

    return (
        <Form className="search" onSubmit={handleSubmitForm}>
            <Word searchRef={searchRef} />
            <Event />
            <Faculty
                faculty={faculty}
                setFaculty={setFaculty}
            />
            <Department
                departmentRef={departmentRef}
                faculty={faculty}
                defaultValue={"Tất cả bộ môn"}
            />
            <Status statusRef={statusRef} />
            <DateFrom dateFromRef={dateFromRef} />
            <DateTo dateToRef={dateToRef} />
            <Form.Group>
                <Button
                    variant="primary"
                    type="submit"
                    className="search-submit"
                >
                    {isLoading ? <Loading></Loading> : <></>}
                    Tìm kiếm
                </Button>
            </Form.Group>
        </Form>
    )

}