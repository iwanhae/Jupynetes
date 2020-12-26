import React from "react";
// react plugin for creating charts
// @material-ui/core
import CardBody from "components/Card/CardBody.js";
import Typography from '@material-ui/core/Typography';

export default function Content(props) {
    return <>
            <Typography component="div" style={{ backgroundColor: '#cfe8fc', height: '22vh' }}>
                <CardBody>
                <p>언제 생성</p>
                <p>몇일째 구동중</p>
                <p>각종 잡다한 세부정보</p>
                <p>구동중인 이미지</p>
                </CardBody>
            </Typography>;
        </>
}