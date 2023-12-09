public class CoordinateBuilder {

  private String left;
  private String right;

  public static CoordinateBuilder builder() {
    return new CoordinateBuilder();
  }

  public CoordinateBuilder withEntry(String unparsedCoordinates) {
    unparsedCoordinates = unparsedCoordinates.replace("(", "").replace(")", "");
    String[] coordinates = unparsedCoordinates.split(", ");
    this.left = coordinates[0];
    this.right = coordinates[1];
    return this;
  }

  public Coordinate build() {
    return new Coordinate(left, right);
  }
}
