
workflow "order-processing" {

    start_event {
    
        service_task "validate_order" {
            # retry {
            #     attempts = 3
            # }
        }
    
    }

    end_event {}

}