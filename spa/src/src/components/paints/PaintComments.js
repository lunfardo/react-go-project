import React, { Component } from 'react'
import { PaintContext, CommentContext } from '../../AppContexts'
import { TextArea } from '../bulmaElements'


class PaintComments extends Component {

    constructor(props) {
        super(props);
        this.state = {
            commentToCreate: '',
        }
    
        this.handleCommentChange = this.handleCommentChange.bind(this);
    }    

    handleCommentChange(event) {
        this.setState({commentToCreate: event.target.value});
    }

    render(){
        return(
        <div>
            <PaintContext.Consumer>{
                paintContext => paintContext.paintsList.filter((paintJSON)=>
                    paintJSON.Id == this.props.match.params.paintId
                ).map((paintJSON) =>
                    <img src={`data:image/png;base64, ${paintJSON.Image}`} />
                )
            }</PaintContext.Consumer>

            <CommentContext.Consumer>{
                commentContext=>
                    <div>
                        <TextArea placeholder="Leave your beautiful comment..." value={this.state.commentToCreate} onChange={this.handleCommentChange} />
                        <a class="button is-info" onClick={commentContext.alert}>Post Comment!</a>
                    </div>
            }</CommentContext.Consumer>                    
        </div>            
        )
    }
}

export default PaintComments