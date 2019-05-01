require 'csv'

CSV.foreach('db/test.csv') do |row|
    Information.create(
        :branch_office => row[0],
        :track_maintenance_area => row[1],
        :control_room => row[2],
        :road_manager_name => row[3],
        :road_name => row[4],
        :location => row[5]
    )
end