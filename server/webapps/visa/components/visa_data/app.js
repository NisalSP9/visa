$(document).ready(function() {
    var visatable = $('#example').DataTable( {
        responsive: true,
        orderCellsTop: true,
        autoWidth: false,
        serverSide: false,
        processing: true,
        lengthMenu: [[10, 25, 50, -1], [10, 25, 50, "All"]],
        language: {
            processing: "<div class='sk-spinner sk-spinner-three-bounce'><div class='sk-bounce1'></div><div class='sk-bounce2'></div><div class='sk-bounce3'></div></div>"
        },
        ajax: {
            //"url": "/DMC_Management_System/HotelController",//http://localhost:3000/api/visa/get
            "url":"/api/visa/get",
            "type": "GET",
            "error": function (e) {
            },
            "dataSrc": function (d) {
                return d;
            }
        },
        deferRender: true,
        columns: [
            {"data": null, "width": "1%", "visible": true},
            {"data": "visaID", "width": "5%", "visible": true},
            {"data": "visaCountry", "width": "5%"},
            {"data": "visaGroupName", "width": "5%"},
            {"data": "visaGroupNameUpdateTime", "width": "5%","visible": false},
            {"data": "visaType", "width": "4%"},
            {"data": "visaTypeEntryUpdateTime", "width": "5%", "visible": false},
            {"data": "visaTypeEntry", "width": "5%"},
            {"data": "visaAgent", "width": "5%"},
            {"data": "visaSubmitByOther", "width": "5%"}

        ],
        colReorder: true,
        dom: "<'row table-toolbar'<'col-sm-1'<'#b3'>><'col-sm-6'B><'col-sm-1'l> <'col-sm-2'f>> <'row'>" +
        "<'row'<'col-sm-12'tr>>" +
        "<'row'<'col-sm-5'i><'col-sm-7'p>>",
        columnDefs: [
            {
                'targets': 0,
                'checkboxes': {
                    'selectRow': true
                }
            }
        ],
        select: {
            'style': 'multi'
        },
        'order': [[1, 'asc']],
        buttons: [
            {
                text: '<i class="fa fa-plus"></i> New Visa',
                className: 'btn-primary',
                container: '#b3',
                action: function (e, dt, node, config) {
                    $('#create-visa-modal').modal()
                }
            },
            'copyHtml5',
            'excelHtml5',
            'csvHtml5',
            'pdfHtml5'
            // {
            //     extend: 'colvis',
            //     text: '<i class="fa fa-gears"></i>',
            //     fade: true
            // }
        ]
    });
    var settings = '<div class="btn-group btn-group-solid">\
            <button type="button" class="btn btn-default dropdown-toggle" style="margin-right:8px"  data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">\
            Action <span class="caret"></span>\
            </button>\
            <ul class="dropdown-menu">\
             <li><a href="#" id="td-edit">Edit</a></li>\
             <li><a href="#" id="td-delete">Delete</a></li>\
            </ul>\
            </div>';

    $("#b3").append(settings);


    // Handle row deletion
    $("#td-delete").on('click', function () {
        if (visatable.column(0).checkboxes.selected().length > 0) {
            $("#confirm-delete").modal();
        }
    });

    $("#btn-delete").on('click', function () {
        var id = visatable.column(0).checkboxes.selected();
        $.each(id, function (k, v) {
            DeleteVisa(v.visaID);
        });
        $("#confirm-delete").modal('hide');
        visatable.ajax.reload();
    });


    $("#btn-visa-save").on('click',function () {
       var visa = JSON.stringify({
           "visaCountry":$("#country").val(),
           "visaGroupName":$("#visaGroupName").val(),
           "visaType":$("#visaType").val(),
           "visaTypeEntry":$("#visaTypeEntry").val(),
           "visaAgent":$("#visaAgent").val(),
           "visaSubmitByOther":$("#visaSubmitByOther").val(),
           "visaQueAppoinnt":$("#visaQueAppoinnt").val(),
           "visaApprovePeriod":$("#visaApprovePeriod").val(),
           "visaRealPrice":$("#visaRealPrice").val(),
           "visaOfficialAgentCharge":$("#visaOfficialAgentCharge").val(),
           "visaSupplierCharge":$("#visaSupplierCharge").val(),
           "visaSupplier":$("#visaSupplier").val(),
           "visaSellingPrice":$("#visaSellingPrice").val(),
           "visaRemark":$("#visaRemark").val(),
           "visaUpdateDateChar":$("#visaUpdateDateChar").val(),
           "visaUpdateDate":$("#visaUpdateDate").val(),
           "visaUpdateBy":$("#visaUpdateBy").val()
        })

        createVisa(visa)

    });


    function createVisa(visa) {

        return $.ajax({
            url: "/api/visa/post",
            type: "POST",
            contentType: "application/json",
            data: visa,
            error: function (e) {
                console.log(e)
            },
            success: function (response) {
                alert("Created");
            }

        });

    }

    function DeleteVisa(id) {
        return $.ajax({
            url: "/api/visa/delete?visaID=" + id,
            type: "DELETE",
            contentType: "application/json",
            error: function (e) {
                console.log(e)
            },
            success: function (response) {
                //alert("delete");
            }

        });
    }


} );