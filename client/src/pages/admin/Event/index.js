import { useState } from "react";
import Card from "../../../components/Shares/Card";
import CustomeProgress from "../../../components/Shares/CustomeProgress";
import "./style.css";
import data from "../../../components/Shares/CustomeProgress/data";
import SubCard from "../../../components/Shares/Card/SubCard";
import { Accordion, Button } from "react-bootstrap";

const Event = () => {
  // eslint-disable-next-line
  const [currentStage, setCurrentStage] = useState(4);
  const [stageDetail, setStageDetail] = useState(currentStage);
  const [update, setUpdate] = useState(false);

  const stage = data.find((item) => item.id === stageDetail);

  return (
    <div className="admin_event">
      <Card title={"Sự kiện đang diễn ra"}>
        <div className="">
          <span style={{ fontWeight: "bold" }}>Sự kiện:</span> Đại hội nghiên
          cứu khoa học lần 36 HUMG
        </div>
        <div className="event_stage">
          <span style={{ fontWeight: "bold" }}>Các giai đoạn:</span>
          <CustomeProgress
            currentStage={currentStage}
            setCurrentStage={setStageDetail}
          />
        </div>
        <div className="">
          <SubCard title={"Chi tiết giai đoạn:"} style={{ fontSize: "16px" }}>
            <div style={{ height: "200px" }}>
              <div className="row stage_detail">
                <div className="col-3">Giai đoạn: </div>
                <div className="col-6">{stage.name}</div>
                <div className="col-2"></div>
              </div>
              <div className="row stage_detail">
                <div className="col-3">Thời gian thực hiện:</div>
                <div className="col-6">{`Từ ${
                  stage.dateFrom ? stage.dateFrom : "..."
                } đến ${stage.dateTo ? stage.dateTo : "..."}`}</div>
                <div className="col-2"></div>
              </div>
              <div className="row stage_detail">
                <div className="col-3">Tóm tắt nội dung: </div>
                <div className="col-6">{stage.description}</div>
                <div className="col-2"></div>
              </div>
              <div className="row stage_detail">
                <div className="col-3"></div>
                <div className="col-6">
                  {update ? (
                    <input type="file"></input>
                  ) : (
                    <span
                      style={{
                        fontWeight: "bold",
                        color: "blue",
                        textDecoration: "underline",
                        cursor: "pointer",
                      }}
                    >
                      Tải tệp đính kèm
                    </span>
                  )}
                </div>
                <div className="col-2"></div>
              </div>
              {stageDetail > currentStage ? (
                <div className="row stage_detail">
                  <div className="col-3"></div>
                  <div className="col-6">
                    <Button onClick={() => setUpdate(true)} style={{marginTop: "12px"}}>Chỉnh sửa</Button>
                  </div>
                  <div className="col-2"></div>
                </div>
              ) : (
                <></>
              )}
            </div>
          </SubCard>
        </div>
      </Card>
      <Card title={"Sự kiện đã hoàn thành"}>
        <Accordion>
          <Accordion.Header className="add-topic-header">
            Đại hội nghiên cứu khoa học sinh viên lần thứ 35 HUMG
          </Accordion.Header>
          <Accordion.Body>
            <div className="">
              <span style={{ fontWeight: "bold" }}>Sự kiện:</span> Đại hội
              nghiên cứu khoa học lần 35 HUMG
            </div>
            <div className="event_stage">
              <span style={{ fontWeight: "bold" }}>Các giai đoạn:</span>
              <CustomeProgress
                currentStage={6}
                setCurrentStage={setStageDetail}
              />
            </div>
            <div className="">
              <SubCard
                title={"Chi tiết giai đoạn:"}
                style={{ fontSize: "16px" }}
              >
                <div style={{ height: "200px" }}>
                  <div className="row stage_detail">
                    <div className="col-3">Giai đoạn: </div>
                    <div className="col-6">{stage.name}</div>
                    <div className="col-2"></div>
                  </div>
                  <div className="row stage_detail">
                    <div className="col-3">Thời gian thực hiện:</div>
                    <div className="col-6">{`Từ ${
                      stage.dateFrom ? stage.dateFrom : "..."
                    } đến ${stage.dateTo ? stage.dateTo : "..."}`}</div>
                    <div className="col-2"></div>
                  </div>
                  <div className="row stage_detail">
                    <div className="col-3">Tóm tắt nội dung: </div>
                    <div className="col-6">{stage.description}</div>
                    <div className="col-2"></div>
                  </div>
                  <div className="row stage_detail">
                    <div className="col-3"></div>
                    <div
                      className="col-6"
                      style={{
                        fontWeight: "bold",
                        color: "blue",
                        textDecoration: "underline",
                        cursor: "pointer",
                      }}
                    >
                      Tải tệp đính kèm
                    </div>
                    <div className="col-2"></div>
                  </div>
                </div>
              </SubCard>
            </div>
          </Accordion.Body>
        </Accordion>
      </Card>
    </div>
  );
};

export default Event;
