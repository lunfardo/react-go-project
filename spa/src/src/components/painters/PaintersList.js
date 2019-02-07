import React, { Component } from 'react'
import { PainterContext } from '../../AppContexts'
import { Link } from 'react-router-dom'

import './PaintersList.css'


const PainterBox = ({painterid, name})=>
    <div class="box">
        <div class="painter-box content">
            <Link to={`/painter/${painterid}`}>{name}</Link>
        </div>
    </div>

class PaintersList extends Component {

    render(){
        return(
            <div class="painters-box">
                <PainterContext.Consumer>{
                    painterContext => painterContext.paintersList.map((painterJSON) =>   
                            <PainterBox painterid={painterJSON.Id} name={painterJSON.Name}/>
                    )
                }</PainterContext.Consumer>                
            </div>
        )
    }
}

export default PaintersList