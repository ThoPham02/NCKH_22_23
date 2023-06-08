import { useState } from "react";
import SwitchCard from "../../../components/Shares/Card/SwitchCard";
import { ResultSearch } from "../../../components/Shares/Search";

const Result = () => {
  const [switchPage, setSwitchPage] = useState(false)
  return (
    <SwitchCard setSwitchPage={setSwitchPage} switchPage={switchPage}>
      {switchPage
        ?
        <>
          <ResultSearch facultyID={0}/>
        </>
        :
        <>
          <ResultSearch />
        </>
      }
    </SwitchCard>
  )
}

export default Result;