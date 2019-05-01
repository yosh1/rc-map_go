class CreateInformation < ActiveRecord::Migration[5.2]
  def change
    create_table :information do |t|
      t.String :branch_office
      t.String :track_maintenance_area
      t.String :control_room
      t.String :road_manager_name
      t.String :road_name
      t.String :location

      t.timestamps
    end
  end
end
