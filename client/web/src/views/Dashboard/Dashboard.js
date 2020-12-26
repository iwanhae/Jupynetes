import React from "react";
// react plugin for creating charts
// @material-ui/core
import { makeStyles } from "@material-ui/core/styles";
// @material-ui/icons
// core components
import GridItem from "components/Grid/GridItem.js";
import GridContainer from "components/Grid/GridContainer.js";
import Instance from "./components/Instance.js"

import styles from "assets/jss/material-dashboard-react/views/dashboardStyle.js";



export default function Dashboard() {

  return (
    <div>
      <GridContainer>
        <GridItem xs={12} sm={6} md={3}>
          <Instance></Instance>
        </GridItem>
      </GridContainer>
    </div>
  );
}

