import React, { CSSProperties, FC } from 'react'
import { Card, Row, Col, Button} from 'antd';
import { ServerObject, ServerStatus } from '../../../model/ServerObject';
import assert from 'assert';
import './Instance.css';
import appState from '../../../AppState';

const gridStyle: CSSProperties = {
  backgroundColor: 'yellow',
  width: "33%",
  textAlign: "center"
};

const infoStyle: CSSProperties = {
  textAlign: "left",
  margin: "50px",
};

const infoBodyStyle: CSSProperties = {
  backgroundColor: 'rgba(255, 255, 0, 1)'
};

const style: CSSProperties = { background: '#0092ff', padding: '8px 0' };

const statusStyle: CSSProperties={
    height: "25px",
    width: "25px",
    backgroundColor: "#F00",
    borderRadius: "50%",
    display: "inline-block",
}

interface InstanceProps {
  serverData: ServerObject
}


const Instance: FC<InstanceProps> = ({serverData}) => {

    const renderInstanceStatus = (serverData:ServerStatus):string => {
        
        switch(serverData) {
            
            case(ServerStatus.ERROR): {
                return "redCircle";
            }
            case(ServerStatus.INITIALIZING): {
                return "yellowCircle";
            }
            case(ServerStatus.RUNNING): {
                return "greenCircle";
            }
            case(ServerStatus.STOPPED): {
                return "greyCircle";
            }
            default: {
                assert(false);
                return "";
            }
        }
    };



    console.log("status: " + renderInstanceStatus(serverData.getStatus()));

                return <Card.Grid style={gridStyle}>
                    <Card title={serverData.name} >
                        <Card extra={<a>정보수정</a>} style={infoStyle} bodyStyle={infoBodyStyle}>
                            <p>{serverData.getFormmatedCreatedAt()} 에 생성 됨</p>
                            <p>{serverData.getNDaysAgo()} 째 구동중 </p>
                            <p>각종 잡다한 세부정보</p>
                            <p>구동중인 이미지</p>
                        </Card>
                        <Row gutter={16}>
                            <Col className="gutter-row" span={6}>
                                <div style={style}>
                                    <span className={renderInstanceStatus(serverData.getStatus())}></span>
                                </div>
                            </Col>
                            <Col className="gutter-row" span={6}>
                                <div style={style}>
                                    <Button danger onClick={() => appState.deleteServer(serverData)}>삭제</Button>
                                </div>
                            </Col>
                            <Col className="gutter-row" span={6}>
                                <div style={style}>
                                    <Button danger>재시작</Button>
                                </div>
                            </Col>
                            <Col className="gutter-row" span={6}>
                                <div style={style}>
                                    <Button danger>접속</Button>
                                </div>
                            </Col>
                        </Row>
                    </Card>
                </Card.Grid>;
}

export default Instance;