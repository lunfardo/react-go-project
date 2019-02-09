import React, { Component } from 'react'
import { PainterContext } from '../../AppContexts'
import { Link } from 'react-router-dom'

import './PaintersList.css'


const PainterBox = ({Id, Name, Image})=>
    <div class="box painter-box">
        <div class="content">
            <div>
                <Link to={`/painter/${Id}`}>{Name}</Link>
            </div>
            <img  class ="painter-picture" src={`data:image/png;base64, ${Image}`}/>
        </div>
    </div>

class PaintersList extends Component {

    render(){
        return(
            <div>
                <div class="painters-box">
                    <PainterContext.Consumer>{
                        painterContext => painterContext.paintersList.map((painterJSON) =>   
                                <PainterBox {...painterJSON} />
                        )
                    }</PainterContext.Consumer> 
                    <p>
                    </p>
                </div>
                <Link to="/painter/create" class="button is-primary">Add New Painter</Link> 
            </div>              

        )
    }
}

export default PaintersList