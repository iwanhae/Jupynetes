import React, { Component } from 'react'
import { Card, Row, Col, Button} from 'antd';

const gridStyle = {
  backgorundColor: 'yellow',
  width: "33%",
  textAlign: "center"
};

const infoStyle = {
  textAlign: "left",
  margin: "50px",
};

const infoBodyStyle = {
  backgorundColor: 'rgba(255, 255, 0, 1)'
};

const style = { background: '#0092ff', padding: '8px 0' };

const statusStyle={
    height: "25px",
    width: "25px",
    backgroundColor: "#F00",
    borderRadius: "50%",
    display: "inline-block",
}

export default class Instance extends Component {
    render() {
        return (
            <>
                <Card.Grid style={gridStyle}>
                    <Card title="사과서버" >
                        <Card extra={<a>정보수정</a>} style={infoStyle} bodyStyle={infoBodyStyle}>
                            <p>언제생성</p>
                            <p>몇일째 구동중</p>
                            <p>각종 잡다한 세부정보</p>
                            <p>구동중인 이미지</p>
                        </Card>
                        <Row gutter={16}>
                            <Col className="gutter-row" span={6}>
                                <div style={style}>
                                    <span class="dot" style={statusStyle}></span>
                                </div>
                            </Col>
                            <Col className="gutter-row" span={6}>
                                <div style={style}>
                                    <Button danger>삭제</Button>
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
                </Card.Grid>
            </>
        )
    }
}