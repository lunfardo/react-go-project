import React, { Component } from 'react';
import { PaintContext } from '../AppContexts'


class PaintProvider extends Component {
    state = {
      paintsList: [],
      submitPaintForm: this.submitPaintForm
    }

    async loadPaintsFromAPI(){
      return fetch(`//${window.location.host}/api/v1/paints`)
      .then((response) => {
        if(response.ok) {
          return response.json()
        } else {
          return 'not ok'
          }
      })
      .catch(function(error) {
        this.setState({apimessage: 'Fetch failed:' + error.message})
      });  
    }

    componentDidMount() {
        this.loadPaintsFromAPI().then((result)=>{
            this.setState({paintsList: result})    
        })          
    }  
    
    async submitPaintForm() {

      const rawResponse = await fetch(`//${window.location.host}/api/v1/paints`, {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({Name: this.state.name, CurrentLocation: this.state.currentLocation, PainterId: parseInt(this.state.painter)})
      });
      
      if(rawResponse.status == 200) {
        this.setState({redirect:true})
      }      
    }

    render() {
      return (
        <PaintContext.Provider value={this.state}>
          {this.props.children}
        </PaintContext.Provider>
      )}
}

export default PaintProvider