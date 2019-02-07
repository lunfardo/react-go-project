import React, { Component } from 'react'
import { PaintContext } from '../../AppContexts'
import { Link } from 'react-router-dom'

const PaintBox = ({Name, CurrentLocation})=>
    <div class="box">
        <h3>{Name}</h3>
        <span>{CurrentLocation}</span>
    </div>


class PaintsList extends Component {

    render(){
        return(
            <div>
                <PaintContext.Consumer>{
                    paintContext => paintContext.paintsList.filter((paintJSON)=>
                        paintJSON.PainterId == this.props.match.params.painterId
                    ).map((paintJSON) =>   
                           <PaintBox {...paintJSON} />
                    )
                }</PaintContext.Consumer>
                <Link to="/paint/create" class="button is-primary">Add new Painting</Link> 
            </div>
        )
    }
}

export default PaintsList