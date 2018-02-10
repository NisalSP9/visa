$(document).ready(function() {
    $('#example').DataTable( {
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
                    window.location.href = "#/new_hotel";
                }
            },
            'copyHtml5',
            'excelHtml5',
            'csvHtml5',
            'pdfHtml5',
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

} );