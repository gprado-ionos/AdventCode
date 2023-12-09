public class NavigationEntryBuilder {

  private String destination;
  private Coordinate coordinate;

  public static NavigationEntryBuilder builder() {
    return new NavigationEntryBuilder();
  }

  public NavigationEntryBuilder withDestination(String destination) {
    this.destination = destination;
    return this;
  }

  public NavigationEntryBuilder withCoordinate(Coordinate coordinate) {
    this.coordinate = coordinate;
    return this;
  }

  public NavigationEntry build() {
    return new NavigationEntry(destination, coordinate);
  }
}
